package com.xdlr.xdublockchainlab.maskvote.ConnectionToChain;

import org.hyperledger.fabric.gateway.*;

import java.io.IOException;
import java.io.InputStream;
import java.io.Reader;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.InvalidKeyException;
import java.security.PrivateKey;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.Properties;


public class Connection {

    /**
     * @description
     */
    public static Network getNetwork() throws IOException, CertificateException, InvalidKeyException {
        Network network = null;

        Properties properties = new Properties();
        InputStream inputStream = com.xdlr.xdublockchainlab.maskvote.ConnectionToChain.Connection.class.getResourceAsStream("/fabric.config.properties");
        properties.load(inputStream);

        String networkConfigProfile = properties.getProperty("networkConfigPath");
        String channelName = properties.getProperty("channelName");
        String certificatePath = properties.getProperty("certificatePath");
        X509Certificate certificate = readX509Certificate(Paths.get(certificatePath));
        String privateKeyPath = properties.getProperty("privateKeyPath");
        PrivateKey privateKey = getPrivateKey(Paths.get(privateKeyPath));

        Wallet wallet = Wallets.newInMemoryWallet();
        wallet.put("user1", Identities.newX509Identity("Org2MSP", certificate,privateKey));



        Gateway.Builder builder = Gateway.createBuilder()
                                    .identity(wallet,"user1")
                                    .networkConfig(Paths.get(networkConfigProfile));

        try(Gateway gateway = builder.connect()){
            network = gateway.getNetwork(channelName);
        }catch (Exception e){
            e.printStackTrace();
        }finally {
            return network;
        }
    }

    private static X509Certificate readX509Certificate(final Path certificatePath) throws IOException, CertificateException {
        try (Reader certificateReader = Files.newBufferedReader(certificatePath, StandardCharsets.UTF_8)) {
            return Identities.readX509Certificate(certificateReader);
        }
    }

    private static PrivateKey getPrivateKey(final Path privateKeyPath) throws IOException, InvalidKeyException {
        try (Reader privateKeyReader = Files.newBufferedReader(privateKeyPath, StandardCharsets.UTF_8)) {
            return Identities.readPrivateKey(privateKeyReader);
        }
    }


}
