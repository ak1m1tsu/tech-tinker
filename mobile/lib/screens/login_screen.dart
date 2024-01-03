import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:tech_tinker/widgets/login_form.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      backgroundColor: CupertinoColors.white,
      body: Center(
        child: SingleChildScrollView(
          child: LoginForm(),
        ),
      ),
    );
  }
}
