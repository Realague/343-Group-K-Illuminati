using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LoginPayload {

	public LoginPayload(string identifier, string password) {
		this.identifier = identifier;
		this.password = password;
	}

	[SerializeField]
	private string identifier;
	public string Identifier {
		get { return identifier; }
		set { identifier = value; }
	}

	[SerializeField]
	private string password;
	public string Password {
		get { return password; }
		set { password = value; }
	}
}
