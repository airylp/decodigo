package com.decodigo.MedianTwoSortedArrays;

public class MedianTwoSortedArrays {
    
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {

        int pos1 = 0;
        int pos2 = 0;
        int l = nums1.length + nums2.length;

        int val1 = 0;
        int val2 = 0;

        for(int i=0;i<= l / 2;i++){
            val1 = val2;
            if( pos1 < nums1.length){
                if(pos2 < nums2.length){
                    if(nums1[pos1] < nums2[pos2]){
                        val2 = nums1[pos1];
                        pos1++;
                    } else {
                        val2 = nums2[pos2];
                        pos2++;
                    }
                } else {
                    val2 = nums1[pos1];
                    pos1++;
                }
            } else {
                if(pos2 < nums2.length){
                    val2 = nums2[pos2];
                    pos2++;
                }
            }
        }

        if(l % 2 == 1){
            return val2;
        } else {
            return (val1 + val2) / 2.0;
        }
        
    }
    

    public static void main(String args[]) {
        System.out.println("res: " + new MedianTwoSortedArrays().findMedianSortedArrays(new int[]{1,3}, new int[]{2}));
    }
}
