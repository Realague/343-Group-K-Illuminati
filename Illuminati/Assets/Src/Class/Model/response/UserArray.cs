using UnityEngine;
using System.Collections;

public class UserArray {

    [SerializeField]
    private User[] users;
    public User[] Users {
        get { return users; }
        set { users = value; }

    }
}
