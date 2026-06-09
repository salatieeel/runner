import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Set;

final class SignatureValidator {
    private static final Set<String> ALLOWED_FORMATS = Set.of("CMS", "XMLDSIG", "PADES");

    void validate(SignatureRequest request) {
        requireNotBlank(request.inputFile(), "inputFile");
        requireExistingFile(request.inputFile(), "inputFile");
        requireNotBlank(request.certificateFile(), "certificateFile");
        requireExistingFile(request.certificateFile(), "certificateFile");

        if ("validate".equals(request.operation())) {
            requireNotBlank(request.signatureFile(), "signatureFile");
            requireExistingFile(request.signatureFile(), "signatureFile");
        }

        if (request.pkcs11Library() != null && !request.pkcs11Library().isBlank()) {
            requireExistingFile(request.pkcs11Library(), "pkcs11Library");
            requireNotBlank(request.pin(), "pin");
        }

        if (request.signatureFormat() != null && !ALLOWED_FORMATS.contains(request.signatureFormat())) {
            throw new ValidationException("signatureFormat invalido. Valores aceitos: " + ALLOWED_FORMATS);
        }
    }

    private void requireNotBlank(String value, String field) {
        if (value == null || value.isBlank()) {
            throw new ValidationException("Campo obrigatorio ausente: " + field);
        }
    }

    private void requireExistingFile(String fileName, String field) {
        if (!Files.exists(Path.of(fileName))) {
            throw new ValidationException("Arquivo informado em " + field + " nao existe: " + fileName);
        }
    }
}
