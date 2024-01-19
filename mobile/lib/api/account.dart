import 'dart:convert';

import 'package:http/http.dart';
import 'package:tech_tinker/api/auth.dart';
import 'package:tech_tinker/constants.dart';
import 'package:tech_tinker/models/account.dart';

class AccountService {
  static var client = Client();

  static Future<Account?> info() async {
    var loginDetails = await AuthCache.loginDetails();
    if (loginDetails == null) {
      return null;
    }

    var url = Uri.http(apiAddress, accountEndpoint);
    var headers = {
      'Authorization': 'Bearer ${loginDetails.token}',
    };

    var resp = await client.get(url, headers: headers);

    if (resp.statusCode == 200) {
      var body = json.decode(resp.body);
      var account = Account.fromJson(body['data']);
      return account;
    }

    return null;
  }
}
