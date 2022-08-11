package com.decodigo.RotacionCiclica;

class RotacionCiclica {
    public int[] solution(int[] A, int K) {
        if(A.length > 1 && K > 0){
            for(int i=0;i<K;i++){
                int aux = A[A.length -1];
                for(int j=A.length -1; j> 0;j--){
                    A[j] = A[j -1]; 
                }
                A[0] = aux;
            }
        }
        return A;
    }

    public static void main(String args[]){
        int []A = {1, 2, 3, 4 ,5};
        RotacionCiclica s = new RotacionCiclica();
        A = s.solution(A, 2);

        for(int i=0;i<A.length;i++){
            System.out.println(": " + A[i]);
        }
    }
}