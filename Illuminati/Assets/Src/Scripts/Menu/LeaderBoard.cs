using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

public class LeaderBoard : MonoBehaviour {

	[SerializeField]
	private TextMeshProUGUI errorText;
	[SerializeField]
	private GameObject ui;
	[SerializeField]
	private GameObject scorePrefab;
	[SerializeField]
	private Transform scoreParent;
	private Player[] players;

	void Start () {
		//StartCoroutine(GetLeaderBoard());
	}

	/*IEnumerator GetLeaderBoard() {
		WWW request = new WWW(RequestUtils.getLeaderBoardURL);

		yield return request;

		if (request.responseHeaders.Count > 0 && request.responseHeaders.ContainsKey("STATUS")) {
            foreach (KeyValuePair<string, string> entry in request.responseHeaders) {
				if (entry.Key == "STATUS" && RequestUtils.ComputeError(entry.Value, errorText)) {
					if (entry.Value.Contains("200")) {
						ui.SetActive(true);
						players = JsonConvert.DeserializeObject<Player[]>(request.text);
						DisplayLeaderBoard();
					}
				} else {
					ui.SetActive(false);
				}
            }
        }
	}*/

	void DisplayLeaderBoard() {
		for (int i = 0; i < players.Length; i++) {
			GameObject score = Instantiate(scorePrefab);
			score.GetComponent<Row>().DisplayInfo(players[i], "#" + (i + 1).ToString());
			score.transform.SetParent(scoreParent);
			score.GetComponent<RectTransform>().localScale = new Vector3(1, 1, 1);
		}
	}

}
