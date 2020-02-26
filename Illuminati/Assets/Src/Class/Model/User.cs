using System.Collections.Generic;
using UnityEngine;

[System.Serializable]
public class User
{
    [SerializeField]
    private string id;
    public string Id {
        get { return id; }
        set { id = value; }
    }

    [SerializeField]
    private string username;
    public string Username {
        get { return username; }
        set { username = value; }
    }

    [SerializeField]
    private string email;
    public string Email {
        get { return email; }
        set { email = value; }
    }

    [SerializeField]
    private List<string> friend_list;
    public List<string> FriendList {
        get { return friend_list; }
        set { friend_list = value; }
    }

    [SerializeField]
    private int mmr;
    public int Mmr {
        get { return mmr; }
        set { mmr = value; }
    }

    [SerializeField]
    private bool verified;
    public bool Verified {
        get { return verified; }
        set { verified = value; }
    }

    [SerializeField]
    private bool admin;
    public bool Admin {
        get { return admin; }
        set { admin = value; }
    }
}
