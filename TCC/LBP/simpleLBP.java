import java.awt.image.BufferedImage;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.PrintWriter;
import java.io.UnsupportedEncodingException;

import javax.imageio.ImageIO;

public class simpleLBP {
	
	public static void main(String[] args){
		try {
			BufferedImage img = ImageIO.read(new File("/home/tiyuri/Imagens/test.jpg"));
			BufferedImage grayVersion = makeGray(img);
			BufferedImage lbpVersion = LBP(grayVersion);
			histogram(lbpVersion);
			System.out.println("terminou");
		} catch (IOException e) {
			e.printStackTrace();
			System.out.println("algo deu errado");
		}
		
		
	}
	
	// Essa função tranforma uma imagem colorida para
	// tons de cinza, mas mantem o formato rgb
	
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
	    
	    File outputfile = new File("/home/tiyuri/Imagens/result.jpg");
	    try {
			ImageIO.write(img, "jpg", outputfile);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			}
	      
	    return(img);
	}
	
	// Essa função calcula o LBP a partir de uma imagem cinza
	// retornando uma nova imagem e um vetor de caracteristicas
	
	public static BufferedImage LBP(BufferedImage img){
		BufferedImage img2 = img;
		int[] neighbours = new int[9];
	    for (int x = 0; x < img.getWidth(); ++x)
	    {
	    	for (int y = 0; y < img.getHeight(); ++y)
	    	{
	    		if(x-1 >= 0 && y-1 >= 0){
	    		neighbours[0] = img.getRGB(x-1, y-1);
	    		neighbours[0] = (neighbours[0] >> 16) & 0xFF;
	    		} else{
	    			neighbours[0] = 0;
	    		}
	    		
	    		if(y-1 >= 0){
	    		neighbours[1] = img.getRGB(x, y-1);
	    		neighbours[1] = (neighbours[1] >> 16) & 0xFF;
	    		}else{
	    			neighbours[1] = 0;
	    		}
	    		
	    		if(x+1 < img.getWidth() && y-1 >= 0){
	    		neighbours[2] = img.getRGB(x+1, y-1);
	    		neighbours[2] = (neighbours[2] >> 16) & 0xFF;
	    		} else{
	    			neighbours[2] = 0;
	    		}
	    		
	    		if(x-1 >= 0){
	    		neighbours[3] = img.getRGB(x-1, y);
	    		neighbours[3] = (neighbours[3] >> 16) & 0xFF;
	    		}else{
	    			neighbours[3] = 0;
	    		}
	    		
	    		if(x+1 < img.getWidth()){
	    		neighbours[4] = img.getRGB(x+1, y);
	    		neighbours[4] = (neighbours[4] >> 16) & 0xFF;
	    		} else{
	    			neighbours[4] = 0;
	    		}
	    		
	    		if(x-1 >= 0 && y+1 < img.getHeight()){
	    		neighbours[5] = img.getRGB(x-1, y+1);
	    		neighbours[5] = (neighbours[5] >> 16) & 0xFF;
	    		}else{
	    			neighbours[5] = 0;
	    		}
	    		
	    		if(y+1 < img.getHeight()){
	    		neighbours[6] = img.getRGB(x, y+1);
	    		neighbours[6] = (neighbours[6] >> 16) & 0xFF;
	    		}else{
	    			neighbours[6] = 0;
	    		}
	    		
	    		if(x+1 < img.getWidth() && y+1 < img.getHeight()){
	    		neighbours[7] = img.getRGB(x+1, y+1);
	    		neighbours[7] = (neighbours[7] >> 16) & 0xFF;
	    		}else{
	    			neighbours[7] = 0;
	    		}
	    		
	    		neighbours[8] = img.getRGB(x, y);
	    		neighbours[8] = (neighbours[8] >> 16) & 0xFF;
	    		
	    		String binnary = "";
	    		for (int i = 0; i < 8; i++){
	    			if(neighbours[i] >= neighbours[8]){
	    				binnary = binnary + "1";
	    			} else{
	    				binnary = binnary + "0";
	    			}
	    		}
	    		int grayLevel = Integer.parseInt(binnary, 2);
	    		//System.out.println(binnary + "\t" + grayLevel);
	    		int newGray = (grayLevel << 16) + (grayLevel << 8) + grayLevel;
	    		img.setRGB(x, y, newGray);
	    		
	    	}
		}
	    
	    File outputfile = new File("/home/tiyuri/Imagens/resultLbp.jpg");
	    try {
			ImageIO.write(img2, "jpg", outputfile);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			}
	    
	    return(img2);
	}
	
	public static void histogram(BufferedImage img){
		int[] histo = new int[256];
		for(int i = 0; i < 256; i++){
			histo[i] = 0;
		}
		
	    for (int x = 0; x < img.getWidth(); ++x)
	    {
	    	for (int y = 0; y < img.getHeight(); ++y)
	    	{
	    		int rgb = img.getRGB(x, y);
	    		int r = (rgb >> 16) & 0xFF;
	    		histo[r] = histo[r] + 1;
	    	}
		}
	    
		try {
			PrintWriter writer = new PrintWriter("/home/tiyuri/Imagens/histogram.txt", "UTF-8");
			for(int i = 0; i < 256; i++){
				writer.print(histo[i] + "\t");
			}
			writer.close();
		} catch (FileNotFoundException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (UnsupportedEncodingException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}

}
