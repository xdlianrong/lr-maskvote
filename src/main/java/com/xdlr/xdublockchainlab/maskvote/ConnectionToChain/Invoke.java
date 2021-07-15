package com.xdlr.xdublockchainlab.maskvote.ConnectionToChain;

import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.sdk.Peer;

import java.util.EnumSet;

public class Invoke {
    public static void invoke(Network network, String str, String newTrans){
        byte[] invokeResult = new byte[0];
        Contract contract = network.getContract("mycc");
        try {
            invokeResult = contract.createTransaction("set")
                    .setEndorsingPeers(network.getChannel().getPeers(EnumSet.of(Peer.PeerRole.ENDORSING_PEER)))
                    .submit(str, newTrans);
            System.out.println(str + "  :已成功上链");
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
