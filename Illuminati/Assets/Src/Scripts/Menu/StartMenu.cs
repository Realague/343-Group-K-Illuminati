using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;

public class StartMenu : MonoBehaviour {

	[SerializeField]
	private TMP_Text welcomeText;

	void Start () {
		welcomeText.text = "WELCOME " + NetControler.instance.player.Username.ToUpper();
	}
	
}
