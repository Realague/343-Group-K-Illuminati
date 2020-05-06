using System.Collections;
using System.Collections.Generic;
using UnityEngine;

[RequireComponent(typeof(AudioSource))]
public class PlayRandomClips : MonoBehaviour
{
    [SerializeField]
    private Object[] clipsList;
    private AudioSource source;
    public static PlayRandomClips instance = null;

    private void Awake() {
        source = GetComponent<AudioSource>();
        //load all the music in the folder specified in parameter\\
        source.clip = clipsList[0] as AudioClip;
        DontDestroyOnLoad(gameObject);
    }

    private void Start() {
        if (instance == null) {
            instance = this;
        } else if (instance != this) {
            Destroy(this.gameObject);
        }
        source.Play();	
	}
    
	private void Update()
    {
		if (!source.isPlaying) {
            playRandomClip();
        }
	}

    private void playRandomClip() {
        source.clip = clipsList[Random.Range(0, clipsList.Length)] as AudioClip;
        source.Play();
    }
}
