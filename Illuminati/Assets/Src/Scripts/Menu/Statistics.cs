using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;

public class Statistics : MonoBehaviour {

	[SerializeField]
	private TMP_InputField usernameInput;
	[SerializeField]
	private TextMeshProUGUI errorText;
	[SerializeField]
	private GameObject info;

	[SerializeField]
	private TextMeshProUGUI username;
	[SerializeField]
	private TextMeshProUGUI winrate;
	[SerializeField]
	private TextMeshProUGUI gamePlayed;
	[SerializeField]
	private TextMeshProUGUI precision;
	[SerializeField]
	private TextMeshProUGUI baseDestroyed;
	[SerializeField]
	private TextMeshProUGUI baseLost;
	[SerializeField]
	private TextMeshProUGUI interception;
	[SerializeField]
	private TextMeshProUGUI score;

	void Start () {
		//DisplayStats(NetControler.instance.player);
	}
	
	void Update () {
		if (Input.GetKeyDown(KeyCode.Return) && usernameInput.text != "") {
			//StartCoroutine(GetStatistics());
		}
	}

	/*IEnumerator GetStatistics() {
		WWW request = new WWW(RequestUtils.getUserStatsURL + usernameInput.text);

		yield return request;

		if (request.responseHeaders.Count > 0 && request.responseHeaders.ContainsKey("STATUS")) {
            foreach (KeyValuePair<string, string> entry in request.responseHeaders) {
				if (entry.Key == "STATUS" && RequestUtils.ComputeError(entry.Value, errorText)) {
					if (entry.Value.Contains("200")) {
						Player player = JsonUtility.FromJson<Player>(request.text);
						DisplayStats(player);
					} else {
						info.SetActive(false);
						RequestError error = JsonUtility.FromJson<RequestError>(request.text);
						errorText.text = error.message;
					}
				} else {
					info.SetActive(false);
				}
            }
			usernameInput.text = "";
        }
	}*/

	void DisplayStats(Player player) {
		double winrateValue = 0;
		if (player.win + player.loose != 0) {
			winrateValue = player.win / (double)(player.win + player.loose) * 100;
		}
		double precisionValue = 0;
		if (player.hit + player.miss != 0) {
			precisionValue = player.hit / (double)(player.hit + player.miss) * 100;
		}
		username.text = player.username;
		winrate.text = (int)winrateValue + "%";
		gamePlayed.text = (player.win + player.loose).ToString();
		precision.text = (int)precisionValue + "%";;
		baseDestroyed.text = player.hit.ToString();
		baseLost.text = player.baseLost.ToString();
		interception.text = player.intercept.ToString();
		score.text = player.mmr.ToString();
		info.SetActive(true);
	}

}
