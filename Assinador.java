import java.io.*;
import java.nio.file.*;
import java.security.MessageDigest;

public class Assinador {

    public static void main(String[] args) {
        try {
            validarArgumentos(args);

            String comando = args[0];

            if (comando.equalsIgnoreCase("assinar")) {
                executarAssinatura(args);
            } else if (comando.equalsIgnoreCase("verificar")) {
                executarVerificacao(args);
            } else {
                erro("Comando inválido. Use: assinar ou verificar.");
            }

        } catch (Exception e) {
            erro("Erro: " + e.getMessage());
        }
    }

    // ===============================
    // ASSINAR
    // ===============================
    private static void executarAssinatura(String[] args) throws Exception {
        String arquivo = getParametro(args, "--arquivo");
        String saida = getParametro(args, "--saida");

        byte[] conteudo = Files.readAllBytes(Paths.get(arquivo));
        String hash = gerarHash(conteudo);

        Files.write(Paths.get(saida), hash.getBytes());

        System.out.println("Arquivo assinado com sucesso!");
        System.out.println("Assinatura gerada: " + hash);
    }

    // ===============================
    // VERIFICAR
    // ===============================
    private static void executarVerificacao(String[] args) throws Exception {
        String arquivo = getParametro(args, "--arquivo");
        String assinaturaArquivo = getParametro(args, "--assinatura");

        byte[] conteudo = Files.readAllBytes(Paths.get(arquivo));
        String hashAtual = gerarHash(conteudo);

        String hashSalvo = new String(Files.readAllBytes(Paths.get(assinaturaArquivo))).trim();

        if (hashAtual.equals(hashSalvo)) {
            System.out.println("Assinatura válida ✅");
        } else {
            System.out.println("Assinatura inválida ❌");
        }
    }

    // ===============================
    // UTILITÁRIOS
    // ===============================
    private static String gerarHash(byte[] dados) throws Exception {
        MessageDigest digest = MessageDigest.getInstance("SHA-256");
        byte[] hashBytes = digest.digest(dados);

        StringBuilder hex = new StringBuilder();
        for (byte b : hashBytes) {
            hex.append(String.format("%02x", b));
        }
        return hex.toString();
    }

    private static void validarArgumentos(String[] args) {
        if (args.length < 3) {
            erro("Uso inválido.\nExemplo:\njava -jar assinador.jar assinar --arquivo a.txt --saida b.txt");
        }
    }

    private static String getParametro(String[] args, String nome) {
        for (int i = 0; i < args.length - 1; i++) {
            if (args[i].equalsIgnoreCase(nome)) {
                return args[i + 1];
            }
        }
        erro("Parâmetro obrigatório não encontrado: " + nome);
        return null;
    }

    private static void erro(String mensagem) {
        System.err.println(mensagem);
        System.exit(1);
    }
}
