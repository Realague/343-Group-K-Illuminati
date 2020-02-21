using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Networking;

public static class ApiClient {

	/*public static void computeError(string value, TextMeshProUGUI errorText) {
		if (value.Contains("50")) {
			errorText.text = "Server Error.";
		} else if (value.Contains("40")) {
			errorText.text = "Connection failed.";
		}
	}*/

	public static T post<T>(string url, byte[] body, string token)
		where T : class {
		using (UnityWebRequest request = new UnityWebRequest(url, "POST")) {
			request.SetRequestHeader("Content-Type", "application/json");
			request.uploadHandler = (UploadHandler)new UploadHandlerRaw(body);
			request.downloadHandler = (DownloadHandler)new DownloadHandlerBuffer();
			addBearerToken(token, request);
			return performRequest<T>(request);
		}
	}

	public static T get<T>(string url, string token)
	where T : class {
		using (UnityWebRequest request = new UnityWebRequest(url, "GET")) {
			request.downloadHandler = (DownloadHandler)new DownloadHandlerBuffer();
			addBearerToken(token, request);
			return performRequest<T>(request);
		}
	}

	public static T put<T>(string url, byte[] body, string token)
	where T : class {
		using (UnityWebRequest request = new UnityWebRequest(url, "PUT")) {
			request.SetRequestHeader("Content-Type", "application/json");
			request.uploadHandler = (UploadHandler)new UploadHandlerRaw(body);
			request.downloadHandler = (DownloadHandler)new DownloadHandlerBuffer();
			addBearerToken(token, request);
			return performRequest<T>(request);
		}
	}

	public static RequestError delete<T>(string url, string token = "M8uqVtkmHWAV3K2PaSZYLKkHWqeCWd22cxGNPXYnpqeT3US")
	where T : class {
		using (UnityWebRequest request = new UnityWebRequest(url, "DELETE")) {
			request.downloadHandler = (DownloadHandler)new DownloadHandlerBuffer();
			addBearerToken(token, request);
			return performRequest<RequestError>(request);
		}
	}

	private static T performRequest<T>(UnityWebRequest request) where T : class {
		request.SendWebRequest();
		while (!request.isDone) ;

		if (request.isNetworkError || request.isHttpError) {
			if (request.downloadHandler.text == "") {
				Debug.Log(request.error);
				//errorText.text = error;
				return null;
			}
			RequestError error = JsonUtility.FromJson<RequestError>(request.downloadHandler.text);
			Debug.Log(error.Error);
			//errorText.text = error.Message;
			return null;
		} else {
			Debug.Log(request.downloadHandler.text);
			return JsonUtility.FromJson<T>(request.downloadHandler.text);
		}
	}

	public static LoginResponse login(string identifier, string encryptedPassword/*, TextMeshProUGUI errorText*/) {
		return null;
	}

	public static string Md5Sum(string strToEncrypt) {
		System.Text.UTF8Encoding ue = new System.Text.UTF8Encoding();
		byte[] bytes = ue.GetBytes(strToEncrypt);

		// encrypt bytes
		System.Security.Cryptography.MD5CryptoServiceProvider md5 = new System.Security.Cryptography.MD5CryptoServiceProvider();
		byte[] hashBytes = md5.ComputeHash(bytes);

		// Convert the encrypted bytes back to a string (base 16)
		string hashString = "";

		for (int i = 0; i < hashBytes.Length; i++) {
			hashString += System.Convert.ToString(hashBytes[i], 16).PadLeft(2, '0');
		}

		return hashString.PadLeft(32, '0');
	}

	private static void addBearerToken(string bearerToken, UnityWebRequest request) {
		request.SetRequestHeader("Authorization", "Bearer " + bearerToken);
	}

	public static byte[] objectToByteArray<T>(T payload)
		where T : class {
		string json = JsonUtility.ToJson(payload);
		Debug.Log(json);
		return System.Text.Encoding.UTF8.GetBytes(json);
	}

}
