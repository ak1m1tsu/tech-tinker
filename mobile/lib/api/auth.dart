import 'dart:convert';

import 'package:api_cache_manager/api_cache_manager.dart';
import 'package:api_cache_manager/models/cache_db_model.dart';
import 'package:http/http.dart' as http;
import 'package:tech_tinker/constants.dart';

class AuthService {
  static var client = http.Client();

  static Future<bool> login(String email, String password) async {
    var url = Uri.http(apiAddress, tokenEndpoint);
    var headers = {
      'Content-Type': 'application/json',
    };

    var resp = await client.post(
      url,
      headers: headers,
      body: jsonEncode({
        "email": email,
        "password": password,
      }),
    );

    if (resp.statusCode == 200) {
      return await AuthCache.setLoginDetails(resp.body);
    }

    return false;
  }
}

class AuthCache {
  static String loginDetailsKey = "login_details";

  static Future<bool> isLoggedIn() async {
    return await APICacheManager().isAPICacheKeyExist(loginDetailsKey);
  }

  static Future<String?> loginDetails() async {
    if (await AuthCache.isLoggedIn()) {
      var cacheData = await APICacheManager().getCacheData(loginDetailsKey);

      return cacheData.syncData;
    }

    return null;
  }

  static Future<bool> setLoginDetails(String? data) async {
    var cache = APICacheDBModel(
      key: loginDetailsKey,
      syncData: data!,
    );

    return await APICacheManager().addCacheData(cache);
  }

  static Future<bool> logout() async {
    return await APICacheManager().deleteCache(loginDetailsKey);
  }
}
