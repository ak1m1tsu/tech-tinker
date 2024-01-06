import 'package:flutter/material.dart';

class Screen extends StatelessWidget {
  final AppBar? appBar;
  final List<Widget> children;

  const Screen({
    super.key,
    required this.children,
    this.appBar,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: appBar,
      body: Padding(
        padding: const EdgeInsets.all(20),
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: children,
          ),
        ),
      ),
    );
  }
}
