import java.util.Map;

final class SignatureApplicationService {
    private final SignatureValidator validator = new SignatureValidator();
    private final FakeSignatureService fakeSignatureService = new FakeSignatureService();
    private final Pkcs11Service pkcs11Service = new Pkcs11Service();

    Map<String, Object> execute(SignatureRequest request) {
        validator.validate(request);
        if ("sign".equals(request.operation())) {
            if (request.pkcs11Library() != null && !request.pkcs11Library().isBlank()) {
                return pkcs11Service.sign(request);
            }
            return fakeSignatureService.sign(request);
        }
        if ("validate".equals(request.operation())) {
            return fakeSignatureService.validate(request);
        }
        throw new ValidationException("Operacao nao suportada: " + request.operation());
    }
}
