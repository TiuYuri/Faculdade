package periocularRecognition;

import java.awt.image.BufferedImage;
import java.io.File;
import java.util.ArrayList;
import java.util.List;

import javax.imageio.ImageIO;

public class Auxiliar {
	public static void main(String[] args) throws Exception{
		String finalPath = "/home/tiyuri/Documentos/tcc/fotosTeste";
		


		BufferedImage img1 = ImageIO.read(new File(finalPath + "/loezer/" + "loezer1.jpg"));
		BufferedImage img2 = ImageIO.read(new File(finalPath + "/loezer/" + "loezer2.jpg"));
		BufferedImage img3 = ImageIO.read(new File(finalPath + "/loezer/" + "loezer3.jpg"));
		BufferedImage img4 = ImageIO.read(new File(finalPath + "/rodrigo/" + "rodrigo1.jpg"));
		BufferedImage img5 = ImageIO.read(new File(finalPath + "/rodrigo/" + "rodrigo2.jpg"));
		BufferedImage img6 = ImageIO.read(new File(finalPath + "/rodrigo/" + "rodrigo3.jpg"));
		BufferedImage img7 = ImageIO.read(new File(finalPath + "/murilo/" + "murilo1.jpg"));
		BufferedImage img8 = ImageIO.read(new File(finalPath + "/murilo/" + "murilo2.jpg"));
		BufferedImage img9 = ImageIO.read(new File(finalPath + "/murilo/" + "murilo3.jpg"));
		BufferedImage img10 = ImageIO.read(new File(finalPath + "/yuri/" + "yuri1.jpg"));
		BufferedImage img11 = ImageIO.read(new File(finalPath + "/yuri/" + "yuri2.jpg"));
		BufferedImage img12 = ImageIO.read(new File(finalPath + "/yuri/" + "yuri3.jpg"));
		BufferedImage img13 = ImageIO.read(new File(finalPath + "/rodrigo/" + "rodrigo4.jpg"));

		
		
		
		User loezer = new User(1, "loezer", 
								FaceRecognition.detectPeriocular(img1),
								FaceRecognition.detectPeriocular(img2),
								FaceRecognition.detectPeriocular(img3)
								);
		
		User salles = new User(2, "salles", 
				FaceRecognition.detectPeriocular(img4),
				FaceRecognition.detectPeriocular(img5),
				FaceRecognition.detectPeriocular(img6)
				);
		
		User murilo = new User(3, "murilo", 
				FaceRecognition.detectPeriocular(img7),
				FaceRecognition.detectPeriocular(img8),
				FaceRecognition.detectPeriocular(img9)
				);
		
		User yuri = new User(4, "yuri", 
				FaceRecognition.detectPeriocular(img10),
				FaceRecognition.detectPeriocular(img11),
				FaceRecognition.detectPeriocular(img12)
				);
		List<User> users = new ArrayList<User>();
		users.add(loezer);
		users.add(salles);
		users.add(murilo);
		users.add(yuri);
		
		
		
		float[] histogram = FaceRecognition.detectPeriocular(img13);		
		BestUser answer = Validator.userListValidate(users, histogram);
		
		System.out.println("Pessoa Escolhida: " + answer.name + " proximidade: " + answer.bestResult );
		
/*
		System.out.println(Validator.personValidate(loezer, histogram));
		System.out.println(Validator.personValidate(salles, histogram));
		System.out.println(Validator.personValidate(murilo, histogram));
		System.out.println(Validator.personValidate(yuri, histogram));	*/	
		
	}
	
}
