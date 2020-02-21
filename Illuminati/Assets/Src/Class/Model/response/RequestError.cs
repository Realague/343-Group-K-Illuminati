using UnityEngine;

public class RequestError
{
	[SerializeField]
	private string error;
	public string Error {
		get { return error; }
		set { error = value; }
	}

}
