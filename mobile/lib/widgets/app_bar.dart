import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

techTinkerAppBar(BuildContext context, String title) {
  return AppBar(
    backgroundColor: CupertinoColors.darkBackgroundGray,
    leading: IconButton(
      icon: const Icon(
        Icons.arrow_back,
        color: CupertinoColors.extraLightBackgroundGray,
      ),
      onPressed: () {
        Navigator.pop(context);
      },
    ),
    title: Text(
      title,
      style: const TextStyle(
        fontWeight: FontWeight.bold,
        fontSize: 24,
        color: CupertinoColors.extraLightBackgroundGray,
      ),
    ),
  );
}
