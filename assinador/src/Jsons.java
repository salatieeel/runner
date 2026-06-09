import java.util.Map;
import java.util.stream.Collectors;

final class Jsons {
    private Jsons() {
    }

    static String stringify(Map<String, ?> values) {
        return values.entrySet()
            .stream()
            .map(entry -> "\"" + escape(entry.getKey()) + "\":" + toJsonValue(entry.getValue()))
            .collect(Collectors.joining(",", "{", "}"));
    }

    private static String toJsonValue(Object value) {
        if (value == null) {
            return "null";
        }
        if (value instanceof Number || value instanceof Boolean) {
            return value.toString();
        }
        return "\"" + escape(value.toString()) + "\"";
    }

    private static String escape(String value) {
        return value.replace("\\", "\\\\").replace("\"", "\\\"");
    }
}
