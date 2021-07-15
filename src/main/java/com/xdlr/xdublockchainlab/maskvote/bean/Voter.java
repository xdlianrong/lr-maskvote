package com.xdlr.xdublockchainlab.maskvote.bean;

import org.springframework.stereotype.Component;

@Component
public class Voter {
   //有一系列选票
    private Integer id;
    private Ballot ballot;
    private String privateKey;

    public Voter(){
        ballot = new Ballot();
    }

    public Voter(Ballot ballot){
        this.ballot = ballot;
    }
}
