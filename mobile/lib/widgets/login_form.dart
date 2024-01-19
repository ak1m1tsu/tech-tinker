import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:tech_tinker/api/auth.dart';

class LoginForm extends StatefulWidget {
  const LoginForm({super.key});

  @override
  State<LoginForm> createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> {
  bool _passwordVisible = false;

  final _formKey = GlobalKey<FormState>();

  final emailController = TextEditingController();
  final passwordController = TextEditingController();

  @override
  void dispose() {
    emailController.dispose();
    passwordController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 15),
            child: TextFormField(
              controller: emailController,
              decoration: const InputDecoration(
                prefixIcon: Icon(
                  Icons.email,
                  color: CupertinoColors.darkBackgroundGray,
                  size: 28,
                ),
                border: UnderlineInputBorder(),
                hintText: "Enter the email address",
                filled: true,
                fillColor: CupertinoColors.extraLightBackgroundGray,
              ),
              validator: (value) {
                if (value != null && value.trim().isEmpty) {
                  return "Email is required";
                }
                return null;
              },
            ),
          ),
          const SizedBox(height: 10),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 15),
            child: TextFormField(
              obscureText: !_passwordVisible,
              enableSuggestions: false,
              autocorrect: false,
              controller: passwordController,
              decoration: InputDecoration(
                prefixIcon: const Icon(
                  Icons.key,
                  color: CupertinoColors.darkBackgroundGray,
                  size: 28,
                ),
                suffixIcon: IconButton(
                  onPressed: () {
                    setState(() {
                      _passwordVisible = !_passwordVisible;
                    });
                  },
                  icon: Icon(
                    _passwordVisible
                        ? CupertinoIcons.eye
                        : CupertinoIcons.eye_slash,
                    color: CupertinoColors.darkBackgroundGray,
                  ),
                ),
                border: const UnderlineInputBorder(),
                hintText: "Enter the password",
                filled: true,
                fillColor: CupertinoColors.extraLightBackgroundGray,
              ),
              validator: (value) {
                if (value != null && value.trim().isEmpty) {
                  return "Password is required";
                }
                return null;
              },
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 20, horizontal: 15),
            child: ElevatedButton(
              style: ElevatedButton.styleFrom(
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(10),
                  side: const BorderSide(color: CupertinoColors.systemGrey),
                ),
                backgroundColor: CupertinoColors.extraLightBackgroundGray,
                minimumSize: const Size.fromHeight(40),
                padding: const EdgeInsets.symmetric(
                  vertical: 15,
                  horizontal: 15,
                ),
                textStyle: const TextStyle(
                  fontWeight: FontWeight.bold,
                  fontSize: 18,
                  color: CupertinoColors.darkBackgroundGray,
                ),
              ),
              onPressed: () {
                if (_formKey.currentState!.validate()) {
                  EasyLoading.show(status: "loading...");

                  AuthService.login(
                    emailController.text.toString(),
                    passwordController.text.toString(),
                  ).then(
                    (response) {
                      EasyLoading.dismiss();

                      if (!response) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          const SnackBar(
                            backgroundColor: CupertinoColors.darkBackgroundGray,
                            content: Text(
                              "Invalid login credentials",
                              style: TextStyle(
                                color: CupertinoColors.systemRed,
                              ),
                            ),
                          ),
                        );
                        return;
                      }

                      Navigator.restorablePushNamedAndRemoveUntil(
                        context,
                        "/home",
                        (route) => false,
                      );
                    },
                  ).catchError((err) {
                    EasyLoading.dismiss();

                    print(err);

                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(
                        content: Text(
                          "Server do not response",
                          style: TextStyle(
                            color: CupertinoColors.systemRed,
                          ),
                        ),
                      ),
                    );
                  });
                }
              },
              child: const Text(
                "Login",
                style: TextStyle(
                  color: CupertinoColors.darkBackgroundGray,
                ),
              ),
            ),
          )
        ],
      ),
    );
  }
}
