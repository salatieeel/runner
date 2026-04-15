package br.ufg.inf.es.runner.assinador;

final class ValidationException extends RuntimeException {
    ValidationException(String message) {
        super(message);
    }
}
