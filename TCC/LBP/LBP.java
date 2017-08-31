package periocularRecognition;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

import javax.imageio.ImageIO;

public class LBP {
	
	public static void main(String[] args){
		try {
			BufferedImage img = ImageIO.read(new File("/home/tiyuri/Documentos/tcc/periocularRecognition/imgs/test.jpg"));
			BufferedImage grayVersion = makeGray(img);
			float[] caracteristics = doLBP(grayVersion);
			
			for(int x = 0; x < 256; ++x){
				System.out.println(x + ": " + caracteristics[x]);
			}
			System.out.println("terminou");
		} catch (IOException e) {
			e.printStackTrace();
			System.out.println("algo deu errado");
		}
		
		
	}
	
	public static float[] returnLbpHistogram(BufferedImage img){
		img = makeGray(img);
		float[] vector = doLBP(img);
		return vector;
	}
	
	/*
	 * transform a rgb image to a gray image
	 */
	public static BufferedImage makeGray(BufferedImage img)
	{
	    for (int x = 0; x < img.getWidth(); ++x)
	    {
	    	for (int y = 0; y < img.getHeight(); ++y)
	    	{
	    		int rgb = img.getRGB(x, y);
	    		int r = (rgb >> 16) & 0xFF;
	    		int g = (rgb >> 8) & 0xFF;
	    		int b = (rgb & 0xFF);

	    		int grayLevel = (r + g + b) / 3;
	    		int gray = (grayLevel << 16) + (grayLevel << 8) + grayLevel;
	    		img.setRGB(x, y, gray);
	    	}
		} 
	    return(img);
	}
	
	// Essa função calcula o LBP a partir de uma imagem cinza
	// retornando uma nova imagem e um vetor de caracteristicas
	
	public static float[] doLBP(BufferedImage img){
		BufferedImage img2 = img;
		String binary = "";
		int mainPixel;
		int[] neighbours = new int[8];
		
		/*
		 * transform the center of image
		 */
		for (int x = 1; x < img.getWidth()-1; ++x){
			for (int y = 1; y < img.getHeight()-1; ++y){
				
				mainPixel = (img.getRGB(x, y) >> 16) & 0xFF;
	    		neighbours[0] = (img.getRGB(x-1, y-1) >> 16) & 0xFF;
	    		neighbours[1] = (img.getRGB(x, y-1) >> 16) & 0xFF;
	    		neighbours[2] = (img.getRGB(x+1, y-1) >> 16) & 0xFF;
	    		neighbours[3] = (img.getRGB(x+1, y) >> 16) & 0xFF;
	    		neighbours[4] = (img.getRGB(x+1, y+1) >> 16) & 0xFF;
	    		neighbours[5] = (img.getRGB(x, y+1) >> 16) & 0xFF;
	    		neighbours[6] = (img.getRGB(x-1, y+1) >> 16) & 0xFF;
	    		neighbours[7] = (img.getRGB(x-1, y) >> 16) & 0xFF;
	    		
	    		for (int i = 0; i < 8; i++){
	    			if(neighbours[i] >= mainPixel){
	    				binary = binary + "1";
	    			} else{
	    				binary = binary + "0";
	    			}
	    		}
	    		int grayLevel = Integer.parseInt(binary, 2);
	    		//System.out.println(binary + "\t" + grayLevel);
	    		int newGray = (grayLevel << 16) + (grayLevel << 8) + grayLevel;
	    		img2.setRGB(x, y, newGray);
	    		binary = "";
			}
		}
		
		/*
		 * transform the first line
		 */
		for (int x = 0; x <img2.getWidth(); ++x){
			img2.setRGB(x, 0, img2.getRGB(x, 1));
		}
		
		/*
		 * transform last line
		 */
		for (int x = 0; x < img2.getWidth(); ++x){
			img2.setRGB(x, img2.getHeight()-1, img2.getRGB(x, img2.getHeight()-2));
		}
		
		/*
		 * transform first column
		 */
		for (int y = 0; y < img2.getHeight(); ++y){
			img2.setRGB(0, y, img2.getRGB(1, y));
		}
		
		/*
		 * transform last column
		 */
		for (int y = 0; y < img2.getHeight(); ++y){
			img2.setRGB(img2.getWidth() -1 , y, img2.getRGB(img2.getWidth()-2, y));
		}
		
		//saveFile(img2, "/home/tiyuri/Documentos/tcc/periocularRecognition/imgs/test2.jpg");
	    
		float[] histogramVector = histogram(img2);
	    return(histogramVector);
	}
	
	
	/*
	 * Receive a image and a file path to save the new image
	 * 
	 */
	public static void saveFile(BufferedImage img, String filePath){
	    File outputfile = new File(filePath);
	    try {
			ImageIO.write(img, "jpg", outputfile);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			}
	}
	
	
	/*
	*this function return the 256 histogram vector
	*
	*/
	public static float[] histogram(BufferedImage img){
		float[] histogram = new float[256];
		for(int i = 0; i < 256; i++){
			histogram[i] = 0;
		}
		
	    for (int x = 0; x < img.getWidth(); ++x)
	    {
	    	for (int y = 0; y < img.getHeight(); ++y)
	    	{
	    		int rgb = img.getRGB(x, y);
	    		int r = (rgb >> 16) & 0xFF;
	    		histogram[r] = histogram[r] + 1;
	    	}
		}
	    
	    for (int x = 0; x < 256; ++x){
	    	histogram[x] = histogram[x]/(img.getHeight()*img.getWidth());
	    }
	    
		return(histogram);
	}

}