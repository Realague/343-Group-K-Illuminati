using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using TMPro;
using System.Text.RegularExpressions;

public class Register : MonoBehaviour {

	[SerializeField]
	private TMP_InputField username;
	[SerializeField]
	private TMP_InputField email;
	[SerializeField]
	private TMP_InputField password;
	[SerializeField]
	private TMP_InputField confirmPassword;
	[SerializeField]
	private TextMeshProUGUI errorText;
	[SerializeField]
	private TextMeshProUGUI successText;
	[SerializeField]
	private GameObject registerMenu;
	[SerializeField]
	private GameObject choiceMenu;
	[SerializeField]
	private GameObject registerButton;
	[SerializeField]
	private GameObject registerButton2;

	void Update () {
		if (Input.GetKeyDown(KeyCode.Tab)) {
			if (username.isFocused) {
				email.Select();
			} else if (email.isFocused) {
				password.Select();
			} else if (password.isFocused) {
				confirmPassword.Select();
			}  else if (confirmPassword.isFocused) {
				username.Select();
			}
		}

		if (Input.GetKeyDown(KeyCode.Return)) {
				RegisterButton();
		}
	}

	public void RegisterButton() {
		errorText.text = "";
		if (username.text != "" && password.text != "" && email.text != "" && confirmPassword.text != "") {
			if (!Regex.Match(email.text, @"^.+@[a-z]+\.[a-z]+$").Success) {
				errorText.text = "Email is not valid.";
				email.text = "";
				return;
			} else if (password.text.CompareTo(confirmPassword.text) != 0) {
				errorText.text = "Passwords are not equal.";
				password.text = "";
				confirmPassword.text = "";
				return;
			} else if (password.text.Length < 6) {
				errorText.text = "Password must contain at least 6 character.";
				password.text = "";
				confirmPassword.text = "";
				return;	
			}

		}

		registerButton2.SetActive(false);
		RequestError responses = IlluminatiService.register(username.text, email.text, password.text);
		if (responses == null || responses.Error == "") {
			successText.text = "Registration successful !\nConfirm your email before loging in.";
			registerButton.SetActive(false);
			choiceMenu.SetActive(true);
			registerMenu.SetActive(false);
		} else {
			errorText.text = responses.Error;
		}
		registerButton2.SetActive(true);
	}

}