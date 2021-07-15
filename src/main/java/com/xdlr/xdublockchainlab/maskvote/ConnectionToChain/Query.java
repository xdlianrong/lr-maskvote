package com.xdlr.xdublockchainlab.maskvote.ConnectionToChain;

import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.ContractException;

import java.io.FileWriter;
import java.io.IOException;
import java.nio.charset.StandardCharsets;

public class Query {
    public static String queryAllPKs(String str, Contract contract){
        String result = null;
        try {
            byte[] queryAllAssets = contract.evaluateTransaction("get",str);
            result = new String(queryAllAssets, StandardCharsets.UTF_8);
        } catch (Exception e) {
            e.printStackTrace();
        }
        return result;
    }

    public static String queryH_indexR_x(String str, Contract contract){
        //链上获取所有计票员h_indexR_x
        String result = null;
        try {
            byte[] queryAllAssets = contract.evaluateTransaction("get",str);
            result = new String(queryAllAssets, StandardCharsets.UTF_8);
        } catch (Exception e) {
            e.printStackTrace();
        }
        return result;
    }

    public static void queryUPK(String str, Contract contract) throws ContractException, IOException {
        byte[] upk = contract.evaluateTransaction("get", str);
        String result = new String(upk, StandardCharsets.UTF_8);

        FileWriter file = new FileWriter("UnionPublicKey.txt");
        file.write(result+"\n");
        file.close();
        System.out.println("联合公钥已成功获取！");
    }

    public static String queryAllVoter(String str, Contract contract){
        //链上获取所有投票员的投票信息
        String result = null;
        try {
            byte[] voteInformation = contract.evaluateTransaction("get",str);
            result = new String(voteInformation, StandardCharsets.UTF_8);
            result = result.replaceAll("\\\\","");
        } catch (Exception e) {
            e.printStackTrace();
        }
        return result;
    }

    public static String queryResult(String str, Contract contract) {
        String result = null;
        try {
            byte[] voteResult = contract.evaluateTransaction("get", str);//这个在做什么还是有点不明白的
            result = new String(voteResult, StandardCharsets.UTF_8);
            System.out.println("此次投票的最后结果："+result+", 请查表！");
        }catch (Exception e){
            e.printStackTrace();
        }finally {
            return result;
        }
    }
}
