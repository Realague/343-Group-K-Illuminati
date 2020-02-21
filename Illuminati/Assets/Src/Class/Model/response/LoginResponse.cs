using UnityEngine;

[System.Serializable]
public class LoginResponse
{
    [SerializeField]
    private string access_token;
    public string AccessToken {
        get { return access_token; }
        set { access_token = value; }
    }

    [SerializeField]
    private string refresh_token;
    public string RefreshToken {
        get { return refresh_token; }
        set { refresh_token = value; }
    }

    [SerializeField]
    private User user;
    public User User {
        get { return user; }
        set { user = value; }
    }
}
