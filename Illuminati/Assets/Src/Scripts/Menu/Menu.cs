using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class Menu : MonoBehaviour {

	private int replay = 0;
	private PhotonView view;

	void Start() {
		view = GetComponent<PhotonView>();
	}

	public void ReturnToMenu() {
		SceneManager.LoadScene("Menu");
	}

	public void StartGame() {
		SceneManager.LoadScene ("Lobby");
	}

	public void QuitGame() {
		Application.Quit ();
	}

	public void Disconnect() {
		if (PhotonNetwork.connected) {
			PhotonNetwork.Disconnect();
		}
		ReturnToMenu();
	}

	public void CloseSession() {
		if (NetControler.instance) {
			Destroy(NetControler.instance.gameObject);
		}
		SceneManager.LoadScene(0);
	}

	public void LoadSceneByName(string name) {
		SceneManager.LoadScene(name);
	}
}
