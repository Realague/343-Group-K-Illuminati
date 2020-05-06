using System.Collections; 
using System.Collections.Generic; 
using UnityEngine; 
using UnityEngine.Audio; 
using UnityEngine.UI;
using TMPro;
 
public class SettingsMenu : MonoBehaviour { 
 
    [SerializeField]
  	private AudioMixer audioMixer;
    [SerializeField]
    private Toggle toggle;
    [SerializeField]
    private TMP_Dropdown resolutionDropdown;

    Resolution[] resolutions;

    void Start() {
        toggle.isOn = Screen.fullScreen;
        resolutions = Screen.resolutions;
        resolutionDropdown.ClearOptions();

        List<string> options = new List<string>();

        int currentIndexResolution = 0;
        for (int i = 0; i < resolutions.Length; i++) {
            string option = resolutions[i].width + " x " + resolutions[i].height;
            options.Add(option);

            if (resolutions[i].width == Screen.currentResolution.width &&
                resolutions[i].height == Screen.currentResolution.height) {
                currentIndexResolution = i;
            }
        }

        resolutionDropdown.AddOptions(options);
        resolutionDropdown.value = currentIndexResolution;
        resolutionDropdown.RefreshShownValue();
    }
  
  	public void SetVolume(float volume) {
        audioMixer.SetFloat("Volume", volume);
    }

    public void SetFullScreen(bool isFullScreen) {
        Screen.fullScreen = isFullScreen;
    }

    public void SetResolution(int ResolutionIndex) {
        Resolution resolution = resolutions[ResolutionIndex];
        Screen.SetResolution(resolution.width, resolution.height, Screen.fullScreen);
    }
 
}
