using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class NetworkTest : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        //RequestError responses = IlluminatiService.register("realague", "julien.delane@epitech.eu", "password");
        LoginResponse response = IlluminatiService.login("realague", "password");
        //Player response2 = IlluminatiService.getByUsername("realague", response.AccessToken);
        //Player response3 = IlluminatiService.addToFriendList(response.User, "5e4c34dabe52e6017877e2d1", response.AccessToken);
        User response4 = IlluminatiService.removeFromFriendList(response.User, "5e4c34dabe52e6017877e2d1", response.AccessToken);
        User[] response5 = IlluminatiService.getRanking(100, 0, response.AccessToken);
    }

    // Update is called once per frame
    void Update()
    {
        
    }
}
