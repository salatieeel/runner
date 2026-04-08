package br.ufg.inf.es.runner.assinador;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.HashMap;
import java.util.Map;

final class SignatureRequestParser {
    SignatureRequest fromCli(CommandLine commandLine, String operation) {
        return new SignatureRequest(
            operation,
            commandLine.option("input-file"),
            commandLine.option("certificate"),
            commandLine.option("signature-file"),
            commandLine.option("pkcs11-lib"),
            commandLine.option("pin"),
            commandLine.option("alias"),
            commandLine.option("signature-format")
        );
    }

    SignatureRequest fromJson(byte[] body, String operation) {
        String payload = new String(body, StandardCharsets.UTF_8).trim();
        Map<String, String> values = new HashMap<>();
        if (payload.startsWith("{") && payload.endsWith("}")) {
            String content = payload.substring(1, payload.length() - 1).trim();
            if (!content.isEmpty()) {
                String[] entries = content.split(",");
                for (String entry : entries) {
                    String[] parts = entry.split(":", 2);
                    if (parts.length == 2) {
                        values.put(unquote(parts[0].trim()), unquote(parts[1].trim()));
                    }
                }
            }
        }

        return new SignatureRequest(
            operation,
            values.get("inputFile"),
            values.get("certificateFile"),
            values.get("signatureFile"),
            values.get("pkcs11Library"),
            values.get("pin"),
            values.get("alias"),
            values.get("signatureFormat")
        );
    }

    private String unquote(String raw) {
        String normalized = raw;
        if (normalized.startsWith("\"") && normalized.endsWith("\"")) {
            normalized = normalized.substring(1, normalized.length() - 1);
        }
        return normalized.replace("\\\"", "\"").replace("\\\\", "\\");
    }
}