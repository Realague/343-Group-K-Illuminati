using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using TMPro;
using UnityEngine.SceneManagement;

public class Login : MonoBehaviour {

	[SerializeField]
	private TMP_InputField username;
	[SerializeField]
	private TMP_InputField password;
	[SerializeField]
	private TMP_Text errorText;
	[SerializeField]
	private GameObject netControllerPrefab;

	void Update () {
		if (Input.GetKeyDown(KeyCode.Tab)) {
			if (username.isFocused) {
				password.Select();
			} else if (password.isFocused) {
				username.Select();
			}
		}

		if (Input.GetKeyDown(KeyCode.Return)) {
				LoginButton();
		}
	}

	public void LoginButton() {
		if (password.text != "" && username.text != "") {
			errorText.text = "";

			LoginResponse response = IlluminatiService.login(username.text, password.text);
			if (response != null) {
				GameObject netController = Instantiate(netControllerPrefab);
				netController.GetComponent<NetControler>().player = response.User;
				SceneManager.LoadScene("Menu");
			} else {
				errorText.text = "Email not comfirmed";
			}
		}
	}

}