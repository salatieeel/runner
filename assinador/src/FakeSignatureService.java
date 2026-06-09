import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.util.LinkedHashMap;
import java.util.Map;

final class FakeSignatureService {
    Map<String, Object> sign(SignatureRequest request) {
        String digest = digest(request.inputFile() + "|" + request.certificateFile() + "|" + defaultValue(request.alias()));
        Map<String, Object> response = new LinkedHashMap<>();
        response.put("operation", "sign");
        response.put("status", "success");
        response.put("signature", "SIMULATED-" + digest);
        response.put("signatureFormat", defaultValue(request.signatureFormat(), "CMS"));
        response.put("message", "Assinatura simulada criada com sucesso.");
        return response;
    }

    Map<String, Object> validate(SignatureRequest request) {
        boolean valid = request.signatureFile() != null && request.signatureFile().toLowerCase().contains("valid");
        Map<String, Object> response = new LinkedHashMap<>();
        response.put("operation", "validate");
        response.put("status", "success");
        response.put("valid", valid);
        response.put("message", valid ? "Assinatura simulada considerada valida." : "Assinatura simulada considerada invalida.");
        return response;
    }

    private String defaultValue(String value) {
        return defaultValue(value, "");
    }

    private String defaultValue(String value, String fallback) {
        return value == null || value.isBlank() ? fallback : value;
    }

    private String digest(String content) {
        try {
            byte[] bytes = MessageDigest.getInstance("SHA-256").digest(content.getBytes(StandardCharsets.UTF_8));
            StringBuilder builder = new StringBuilder();
            for (int index = 0; index < 16; index++) {
                builder.append(String.format("%02x", bytes[index]));
            }
            return builder.toString();
        } catch (Exception error) {
            throw new IllegalStateException("Falha ao gerar assinatura simulada", error);
        }
    }
}
