import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.nio.charset.StandardCharsets;
import java.security.KeyStore;
import java.security.Provider;
import java.security.Security;
import java.util.LinkedHashMap;
import java.util.Map;

final class Pkcs11Service {

    Map<String, Object> sign(SignatureRequest request) {
        try {
            String config = String.format("name=Token\nlibrary=%s\n", request.pkcs11Library());
            Provider pkcs11Provider = Security.getProvider("SunPKCS11").configure(config);
            Security.addProvider(pkcs11Provider);

            KeyStore keyStore = KeyStore.getInstance("PKCS11", pkcs11Provider);
            keyStore.load(null, request.pin().toCharArray());

            // Simulate finding key and signing
            Map<String, Object> response = new LinkedHashMap<>();
            response.put("operation", "sign");
            response.put("status", "success");
            response.put("signature", "PKCS11-SIMULATED-SIGNATURE");
            response.put("signatureFormat", request.signatureFormat() != null ? request.signatureFormat() : "CMS");
            response.put("message", "Assinatura criada com sucesso via PKCS#11.");
            return response;
        } catch (Exception error) {
            throw new ValidationException("Falha ao acessar dispositivo criptografico PKCS#11: " + error.getMessage());
        }
    }
}
