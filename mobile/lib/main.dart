import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:tech_tinker/constants.dart';
import 'package:tech_tinker/screens/home_screen.dart';

Widget _defaultHome = const Scaffold();

void main() {
  WidgetsFlutterBinding.ensureInitialized();

  if (true) {
    _defaultHome = const HomeScreen();
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
        "/": (context) => _defaultHome,
        "/home": (context) => const HomeScreen(),
        "/login": (context) => const Scaffold(),
      },
    );
  }
}
