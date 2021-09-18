package com.guillotine.guillotine;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DelSolController {
    @Autowired
    ContentResolver resolver;


    @PostMapping("/artificialselector")
    String isDeserveSalvation(@RequestParam("masterkey")String masterkey,
                              @RequestParam("mastertail")String mastertail){

        if (masterkey.equals(Inventory.TOKEN)) {
            if(resolver.resolve(mastertail))
            return "Done";
            else return "In process";
        }
        else return  "Token error";
    }
}


