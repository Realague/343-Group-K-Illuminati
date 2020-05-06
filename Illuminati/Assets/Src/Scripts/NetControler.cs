using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class NetControler : MonoBehaviour {

    public User player;
    private string gameVersion = "1.0";

    public static NetControler instance;

    void Start() {
        if (instance == null) {
            instance = this;
        } else {
            Destroy(this.gameObject);
        }
        DontDestroyOnLoad(this.gameObject);
    }

    private void Update() {
    }

}