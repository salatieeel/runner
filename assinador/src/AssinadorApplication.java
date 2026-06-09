import java.io.IOException;
import java.util.Map;

public final class AssinadorApplication {
    private AssinadorApplication() {
    }

    public static void main(String[] args) throws IOException {
        CommandLine commandLine = CommandLine.parse(args);
        String command = commandLine.positional(0);
        if (command == null || "help".equals(command)) {
            printHelp();
            return;
        }

        if ("server".equals(command)) {
            int port = toInt(commandLine.option("port"), 18080);
            int idleTimeout = toInt(commandLine.option("idle-timeout-minutes"), 15);
            new AssinadorHttpServer(port, idleTimeout);
            System.out.println(Jsons.stringify(Map.of("status", "started", "port", port, "idleTimeoutMinutes", idleTimeout)));
            return;
        }

        SignatureApplicationService service = new SignatureApplicationService();
        SignatureRequestParser parser = new SignatureRequestParser();

        try {
            SignatureRequest request = parser.fromCli(commandLine, command);
            System.out.println(Jsons.stringify(service.execute(request)));
        } catch (ValidationException error) {
            System.err.println(Jsons.stringify(Map.of("status", "error", "message", error.getMessage())));
            System.exit(1);
        }
    }

    private static int toInt(String value, int defaultValue) {
        if (value == null || value.isBlank()) {
            return defaultValue;
        }
        return Integer.parseInt(value);
    }

    private static void printHelp() {
        System.out.println("""
            Uso:
              java -jar assinador.jar sign --input-file arquivo --certificate cert.pem
              java -jar assinador.jar validate --input-file arquivo --signature-file assinatura.sig --certificate cert.pem
              java -jar assinador.jar server --port 18080 --idle-timeout-minutes 15
            """);
    }
}
