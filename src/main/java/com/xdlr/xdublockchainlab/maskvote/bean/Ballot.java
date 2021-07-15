package com.xdlr.xdublockchainlab.maskvote.bean;

import lombok.Data;

import java.math.BigInteger;

@Data
public class Ballot {
    private BigInteger firstCom;
    private BigInteger secondCom;


    /**
     * 下面的字段用于证明firstCom中的隐藏值非0即1
     */
    private BigInteger firstComTemp1;
    private BigInteger firstComTemp2;
    private BigInteger hash_challenge;
    private BigInteger firstComResponse1;
    private BigInteger firstComResponse2;
    private BigInteger firstComResponse3;

    /**
     * 下面的字段用于证明公钥的r和secondCom的r相等
     */
    private BigInteger indexEqualT;
    private BigInteger indexEqualS1;
    private BigInteger indexEqualS2;
    private BigInteger indexEqualS3;

    /**
     * 下面的字段用于证明确实是用联合公钥加密的
     */
    private BigInteger PKEncryptTemp;
    private BigInteger PKEncryptHash;
    private BigInteger PKEncryptS1;
    private BigInteger PKEncryptS2;

    public Ballot() {
    }

    public Ballot(BigInteger firstCom, BigInteger secondCom, BigInteger firstComTemp1, BigInteger firstComTemp2, BigInteger hash_challenge,
                BigInteger firstComResponse1, BigInteger firstComResponse2, BigInteger firstComResponse3, BigInteger indexEqualT, BigInteger indexEqualS1,
                BigInteger indexEqualS2, BigInteger indexEqualS3, BigInteger PKEncryptTemp, BigInteger PKEncryptHash, BigInteger PKEncryptS1,
                BigInteger PKEncryptS2) {
        this.firstCom = firstCom;
        this.secondCom = secondCom;
        this.firstComTemp1 = firstComTemp1;
        this.firstComTemp2 = firstComTemp2;
        this.hash_challenge = hash_challenge;
        this.firstComResponse1 = firstComResponse1;
        this.firstComResponse2 = firstComResponse2;
        this.firstComResponse3 = firstComResponse3;
        this.indexEqualT = indexEqualT;
        this.indexEqualS1 = indexEqualS1;
        this.indexEqualS2 = indexEqualS2;
        this.indexEqualS3 = indexEqualS3;
        this.PKEncryptTemp = PKEncryptTemp;
        this.PKEncryptHash = PKEncryptHash;
        this.PKEncryptS1 = PKEncryptS1;
        this.PKEncryptS2 = PKEncryptS2;
    }
}
