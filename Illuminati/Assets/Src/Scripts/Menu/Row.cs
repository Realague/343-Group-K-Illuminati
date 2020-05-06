using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using TMPro;

public class Row : MonoBehaviour {

	[SerializeField]
	private TextMeshProUGUI rank;
	[SerializeField]
	private TextMeshProUGUI username;
	[SerializeField]
	private TextMeshProUGUI score;

	public void DisplayInfo(Player player, string rank) {
		this.rank.text = rank;
		username.text = player.username;
		score.text = player.mmr.ToString();
	}
}
