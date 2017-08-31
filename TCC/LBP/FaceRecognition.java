package periocularRecognition;

import java.awt.image.BufferedImage;


import org.opencv.core.Mat;
import org.opencv.core.MatOfRect;
import org.opencv.core.Point;
import org.opencv.core.Rect;
import org.opencv.core.Scalar;
import org.opencv.imgcodecs.Imgcodecs;
import org.opencv.imgproc.Imgproc;
import org.opencv.objdetect.CascadeClassifier;

class FaceRecognition {
  public static float[] detectEyes(BufferedImage img, String path, String name) {

	// Convert the bufferedImage to a Mat type
	
	Mat image = Converter.BufferedImageToMat(img);
	
	//Loads a cascade classifier to eye recognition
	
	CascadeClassifier faceDetector = new CascadeClassifier("/home/tiyuri/Documentos/tcc/periocularRecognition/code/haarcascade_eye_tree_eyeglasses.xml");
	
	if(!faceDetector.empty()){
	    // Detect faces in the image.
	    // MatOfRect is a special container class for Rect.
	    MatOfRect faceDetections = new MatOfRect();
	    faceDetector.detectMultiScale(image, faceDetections);
	
	    System.out.println(String.format("Detected %s faces", faceDetections.toArray().length));
	    
	    
	    int i = 0;
	    // Draw a bounding box around each face.
	    // Make new images for each box.
	    
	    for (Rect rect : faceDetections.toArray()) {
	    	i ++;  	
	    	String filename = path + name + i + ".png";
	    	System.out.println(filename);
	    	Mat cropped = new Mat(image, rect);
	    	Imgcodecs.imwrite(filename, cropped);
	        Imgproc.rectangle(image, new Point(rect.x, rect.y), new Point(rect.x + rect.width, rect.y + rect.height), new Scalar(0, 255, 0));
	    }
	
	} else{
    	System.out.println("classifier not loaded");
    }
    
    float[] histogram = new float[256];
    return histogram;
  }
  
  public static float[] detectPeriocular(BufferedImage img) throws Exception {

		// Convert the bufferedImage to a Mat type
		
		Mat image = Converter.BufferedImageToMat(img);
		
		//Loads a cascade classifier to eye recognition
		
		CascadeClassifier faceDetector = new CascadeClassifier("/home/tiyuri/Documentos/tcc/periocularRecognition/code/haarcascade_eye_tree_eyeglasses.xml");
		
		
		if(!faceDetector.empty()){
		    // Detect faces in the image.
		    // MatOfRect is a special container class for Rect.
		    MatOfRect faceDetections = new MatOfRect();
		    faceDetector.detectMultiScale(image, faceDetections);
		
		    System.out.println(String.format("Detected %s faces", faceDetections.toArray().length));
		
		    //Get the periocular region of image
		    
		    if(faceDetections.toArray().length == 2){
		    	int x0 = 99999, y0 = 99999, x1 = 0, y1 = 0;
		    	for (Rect rect : faceDetections.toArray()) {
		    		if(rect.x < x0){x0 = rect.x;}
		    		if(rect.y < y0){y0 = rect.y;}
		    		if(x1 < (rect.x + rect.width)){x1 = rect.x + rect.width;}
		    		if(y1 < (rect.y + rect.height)){y1 = rect.y + rect.height;}
			    }
		        Rect rect = new Rect(new Point(x0, y0), new Point(x1, y1));
		        Mat matImage = image.submat(rect);
		        BufferedImage periocular = Converter.MatToBufferedImage(matImage);
			    
		        //Return the LBP result
		        return LBP.returnLbpHistogram(periocular);
		    }
		
		} else{
	    	System.out.println("classifier not loaded");
		}
		
    	float[] a = new float[256];
    	a[0] = 999;
    	return a;
	  }
}