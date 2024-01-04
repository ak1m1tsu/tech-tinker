import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:tech_tinker/constants.dart';
import 'package:tech_tinker/widgets/login_form.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: CupertinoColors.darkBackgroundGray,
        title: const Text(
          appName,
          style: TextStyle(
            color: CupertinoColors.extraLightBackgroundGray,
            fontWeight: FontWeight.bold,
            fontSize: 32,
          ),
        ),
        centerTitle: true,
      ),
      backgroundColor: CupertinoColors.extraLightBackgroundGray,
      body: const Center(
        child: SingleChildScrollView(
          child: LoginForm(),
        ),
      ),
    );
  }
}
