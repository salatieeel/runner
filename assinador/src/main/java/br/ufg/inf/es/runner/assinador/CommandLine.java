package br.ufg.inf.es.runner.assinador;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

final class CommandLine {
    private final List<String> positionals;
    private final Map<String, String> options;

    private CommandLine(List<String> positionals, Map<String, String> options) {
        this.positionals = positionals;
        this.options = options;
    }

    static CommandLine parse(String[] args) {
        List<String> positionals = new ArrayList<>();
        Map<String, String> options = new HashMap<>();

        for (int index = 0; index < args.length; index++) {
            String token = args[index];
            if (!token.startsWith("--")) {
                positionals.add(token);
                continue;
            }

            String normalized = token.substring(2);
            int separator = normalized.indexOf('=');
            if (separator >= 0) {
                options.put(normalized.substring(0, separator), normalized.substring(separator + 1));
                continue;
            }

            if (index + 1 >= args.length || args[index + 1].startsWith("--")) {
                options.put(normalized, "true");
                continue;
            }

            options.put(normalized, args[index + 1]);
            index += 1;
        }

        return new CommandLine(positionals, options);
    }

    String positional(int index) {
        return positionals.size() > index ? positionals.get(index) : null;
    }

    String option(String key) {
        return options.get(key);
    }
}
