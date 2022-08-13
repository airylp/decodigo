package com.decodigo.LongestSubstring;

public class LongestSubstring {

    public int lengthOfLongestSubstring(String s){
        System.out.println("s: " + s);
        int map[] = new int[128];

        int l = 0;
        int start = 0;
        if(s.length() == 1){
            return 1;
        }
        if(s.length() > 1){
            for(int i=0;i<s.length();i++){
                int c = s.charAt(i);
                if(map[c] > 0 && map[c] > start){
                    start = map[c];
                }
                map[c] = i + 1;     //Se suma con uno para que cualquier elemento no esté en el índice 0
                if(i + 1 - start > l ) {
                    System.out.println("start: " + start + " i: " + i); //3
                    l = i + 1 - start ;
                }
            }

            return l;
        }
        return l;
    }
    
    public static void main(String args[]){
        LongestSubstring l = new LongestSubstring();
        
        //l.lengthOfLongestSubstring("awwab");
        //l.lengthOfLongestSubstring("jbpnbwwd");
        System.out.println("-l: " + l.lengthOfLongestSubstring("1")); //1
        System.out.println("-l: " + l.lengthOfLongestSubstring("au")); //2
        System.out.println("-l: " + l.lengthOfLongestSubstring("abba")); //2
        System.out.println("-l: " + l.lengthOfLongestSubstring("jbpnbwwd")); //4
        System.out.println("-l: " + l.lengthOfLongestSubstring("abcabcbb")); //3
        System.out.println("-l: " + l.lengthOfLongestSubstring("bbbbb")); //1
        System.out.println("-l: " + l.lengthOfLongestSubstring("pwwkew")); //3
    }
}