package com.guillotine.guillotine;

import org.apache.tomcat.util.codec.binary.Base64;
import org.springframework.stereotype.Component;

import java.util.Arrays;

@Component
public class ContentResolver {

   boolean resolve(String content){

       byte[] maskOFF = decode64(content);
       for(int i=0; i<maskOFF.length;i++){
           maskOFF[i]+=5;
       }
       if(new String(maskOFF).equals(Inventory.TASKVALIDATIONKEY))
       return true;
       else return false;
   };

   byte[] decode64(String sixtyfour){
      byte[] bytes = Base64.decodeBase64(sixtyfour);
      if (bytes.length>5){
         bytes = Arrays.copyOfRange(bytes,bytes.length-6,bytes.length-1);
      } else return bytes;
      return bytes;
   }
}
