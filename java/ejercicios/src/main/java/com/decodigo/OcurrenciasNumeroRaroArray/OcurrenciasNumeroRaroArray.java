package com.decodigo.OcurrenciasNumeroRaroArray;
        
import java.util.HashMap;
import java.util.Map;

public class OcurrenciasNumeroRaroArray {

    public int solucion(int []a){
        Map<Integer,Integer> m = new HashMap<Integer,Integer>();
        
        for (int i = 0; i < a.length; i++) {
            Integer val = m.get(a[i]);
            if(val == null){
                m.put(a[i], 1);
            } else {
                m.put(a[i], val + 1);
            }
        }
        
        for (Integer k : m.keySet()) {
            if(m.get(k).intValue() % 2 == 1){
                return k.intValue();
            }
        }
        return 0;
    }
    
    public static void main(String args[]) {
        int []a = {1, 1, 2, 2, 2, 6 ,6 , 8, 4, 4 , 8};
        
        OcurrenciasNumeroRaroArray o = new OcurrenciasNumeroRaroArray();
        System.out.println("resultado: "+ o.solucion(a));
        
    }
}
