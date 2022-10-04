from imutils import face_utils
import dlib
import cv2


# Se inicializa primero el detector de rostros (HOG) y después
# se agregan los puntos de referencia en el rostro detectado

# p = Es la referencia al modelo preentrenado, en este caso en
#     la misma carpeta que este script
p = "shape_predictor_68_face_landmarks.dat"
detector = dlib.get_frontal_face_detector()
predictor = dlib.shape_predictor(p)

cap = cv2.VideoCapture(0)
 
while True:
    # Se obtiene la imagen de la webcam
    _, image = cap.read()
    # Converting the image to gray scale
    # Se convierte la imagen a escala de grises
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
        
    # Se obtienen los rostros de la imagen en la cámara web
    rects = detector(gray, 0)
    
    # Para cada rostro detectado se buscan los puntos de referencia
    for (i, rect) in enumerate(rects):
    
        # Se hace la predicción y se transforma en un arreglo numpy
        shape = predictor(gray, rect)
        shape = face_utils.shape_to_np(shape)
    
        # Se dibuja en nuestra imagen, todos las coordenadas encontradas (x,y)
        for (x, y) in shape:
            cv2.circle(image, (x, y), 2, (0, 255, 0), -1)
    
    # Se muestra una ventana con la imagen
    cv2.imshow("Output", image)
    
    # La ventana cierra con la tecla Esc
    k = cv2.waitKey(5) & 0xFF
    if k == 27:
        break

cv2.destroyAllWindows()
cap.release()
