package periocularRecognition;

public class User {
	int _id;
	String nome;
	float[] photo1;
	float[] photo2;
	float[] photo3;
	
	public User(int id, String nome, float[] photo1, float[] photo2, float[] photo3 ){
		this._id = id;
		this.nome = nome;
		this.photo1 = photo1;
		this.photo2 = photo2;
		this.photo3 = photo3;
	}
	
}
