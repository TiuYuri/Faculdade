package periocularRecognition;

import java.util.List;

public class Validator {
	
	//Compare 2 histograms and return the result
	public static float validate(float[] histogram1, float[] histogram2){
		double sum = 0;
		for(int x = 0; x < 256; ++x){
			sum = Math.pow((histogram1[x] - histogram2[x]), 2);
		}
		
		return (float) Math.sqrt(sum);
	}
	
	//compare the histogram of a user and return the best result 
	public static float personValidate(User user, float[] histogram){
		float sum = validate(user.photo1, histogram);
		float sum2 = validate(user.photo2, histogram);
		if(sum2 < sum){sum = sum2;}
		sum2 = validate(user.photo3, histogram);
		if(sum2 < sum){sum = sum2;}
		return sum;
	}
	
	public static BestUser userListValidate(List<User> users, float[] histogram){
		float validate;
		BestUser answer = new BestUser(0,"",999);
		for(User a : users){
			validate = personValidate(a, histogram);
			if (validate < answer.bestResult){
				answer._id = a._id;
				answer.name = a.nome;
				answer.bestResult = validate;
			}
		}
		return answer;
	}
	
}
