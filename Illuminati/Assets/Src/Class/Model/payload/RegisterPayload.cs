using UnityEngine;
using System.Collections;

public class RegisterPayload {

	public RegisterPayload(string username, string email, string password) {
		this.username = username;
		this.email = email;
		this.password = password;
	}

	[SerializeField]
	private string username;
	public string Username {
		get { return username; }
		set { username = value; }
	}

	[SerializeField]
	private string email;
	public string Email {
		get { return email; }
		set { email = value; }
	}

	[SerializeField]
	private string password;
	public string Password {
		get { return password; }
		set { password = value; }
	}

}
