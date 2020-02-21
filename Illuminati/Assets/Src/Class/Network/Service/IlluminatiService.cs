using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public static class IlluminatiService {

	public static string authenticationKey = "M8uqVtkmHWAV3K2PaSZYLKkHWqeCWd22cxGNPXYnpqeT3US";
	public static string apiEndPoint = "http://localhost:3000/api";
	public static string loginUrl = apiEndPoint + "/auth/local";

	public static string updateDataURL = apiEndPoint + "/statistics/update";
	public static string getUserStatsURL = apiEndPoint + "/user/";
	public static string usersEndpoint = apiEndPoint + "/users";
	public static string getLeaderBoardURL = apiEndPoint + "/statistics";

	public static LoginResponse login(string identifier, string password) {
		password = ApiClient.Md5Sum(password);
		LoginPayload loginPayload = new LoginPayload(identifier, password);
		return ApiClient.post<LoginResponse>(loginUrl, ApiClient.objectToByteArray<LoginPayload>(loginPayload), authenticationKey);
	}

	public static RequestError register(string username, string email, string password) {
		password = ApiClient.Md5Sum(password);
		RegisterPayload registerPayload = new RegisterPayload(username, email, password);
		return ApiClient.post<RequestError>(loginUrl, ApiClient.objectToByteArray<RegisterPayload>(registerPayload), authenticationKey);
	}

	public static User[] getRanking(int limit, int page, string token) {
		return ApiClient.get<UserArray>(usersEndpoint + "?limit=" + limit + "&offset=" + page + "&sort_by=-mmr", token).Users;
	}

	public static User getByUsername(string username, string token) {
		return ApiClient.get<UserArray>(usersEndpoint + "?limit=1&q=username=/" + username, token).Users[0];
	}

	public static User addToFriendList(User player, string id, string token) {
		player.FriendList.Add(id);
		return updatePlayer(player, token);
	}

	public static User removeFromFriendList(User player, string id, string token) {
		player.FriendList.Remove(id);
		return updatePlayer(player, token);
	}

	public static User updatePlayer(User player, string token) {
		return ApiClient.put<User>(usersEndpoint + "/" + player.Id, ApiClient.objectToByteArray<User>(player), token);
	}


}
