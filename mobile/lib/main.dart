import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:tech_tinker/api/auth.dart';
import 'package:tech_tinker/constants.dart';
import 'package:tech_tinker/screens/home_screen.dart';
import 'package:tech_tinker/screens/login_screen.dart';

Widget _defaultWelcomeScreen = const LoginScreen();

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  if (await AuthCache.isLoggedIn()) {
    _defaultWelcomeScreen = const HomeScreen();
  }

  runApp(const TechTinkerApp());
}

class TechTinkerApp extends StatelessWidget {
  const TechTinkerApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: appName,
      theme: ThemeData(
        primaryColor: CupertinoColors.systemBlue,
        useMaterial3: true,
      ),
      routes: {
        "/": (context) => _defaultWelcomeScreen,
        "/home": (context) => const HomeScreen(),
        "/login": (context) => const LoginScreen(),
      },
      builder: EasyLoading.init(),
    );
  }
}
