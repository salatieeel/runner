import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.nio.charset.StandardCharsets;
import java.time.Instant;
import java.util.Map;
import java.util.concurrent.Executors;

final class AssinadorHttpServer {
    private final SignatureApplicationService service = new SignatureApplicationService();
    private final SignatureRequestParser parser = new SignatureRequestParser();
    private Instant lastInteraction = Instant.now();
    private final int idleTimeoutMinutes;

    AssinadorHttpServer(int port, int idleTimeoutMinutes) throws IOException {
        this.idleTimeoutMinutes = idleTimeoutMinutes;
        HttpServer server = HttpServer.create(new InetSocketAddress(port), 0);
        server.createContext("/health", exchange -> writeResponse(exchange, 200, Jsons.stringify(Map.of("status", "ok"))));
        server.createContext("/sign", new OperationHandler("sign"));
        server.createContext("/validate", new OperationHandler("validate"));
        server.createContext("/shutdown", exchange -> {
            writeResponse(exchange, 200, Jsons.stringify(Map.of("status", "stopping")));
            new Thread(() -> {
                try {
                    Thread.sleep(200L);
                } catch (InterruptedException ignored) {
                    Thread.currentThread().interrupt();
                }
                server.stop(0);
                System.exit(0);
            }).start();
        });
        server.setExecutor(Executors.newCachedThreadPool());

        Thread idleMonitor = new Thread(() -> {
            while (true) {
                if (Instant.now().isAfter(lastInteraction.plusSeconds(idleTimeoutMinutes * 60L))) {
                    server.stop(0);
                    System.exit(0);
                }
                try {
                    Thread.sleep(15_000L);
                } catch (InterruptedException ignored) {
                    Thread.currentThread().interrupt();
                    return;
                }
            }
        });
        idleMonitor.setDaemon(true);
        idleMonitor.start();
        server.start();
    }

    private final class OperationHandler implements HttpHandler {
        private final String operation;

        private OperationHandler(String operation) {
            this.operation = operation;
        }

        @Override
        public void handle(HttpExchange exchange) throws IOException {
            lastInteraction = Instant.now();
            try {
                SignatureRequest request = parser.fromJson(exchange.getRequestBody().readAllBytes(), operation);
                Map<String, Object> response = service.execute(request);
                writeResponse(exchange, 200, Jsons.stringify(response));
            } catch (ValidationException error) {
                writeResponse(exchange, 400, Jsons.stringify(Map.of("status", "error", "message", error.getMessage())));
            }
        }
    }

    private void writeResponse(HttpExchange exchange, int statusCode, String payload) throws IOException {
        byte[] bytes = payload.getBytes(StandardCharsets.UTF_8);
        exchange.getResponseHeaders().add("Content-Type", "application/json; charset=utf-8");
        exchange.sendResponseHeaders(statusCode, bytes.length);
        try (OutputStream outputStream = exchange.getResponseBody()) {
            outputStream.write(bytes);
        }
    }
}
