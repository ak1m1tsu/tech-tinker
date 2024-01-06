import 'package:flutter/cupertino.dart';
import 'package:tech_tinker/models/configuration.dart';
import 'package:tech_tinker/models/customer.dart';
import 'package:tech_tinker/models/order.dart';
import 'package:tech_tinker/widgets/order_list_view.dart';
import 'package:tech_tinker/widgets/screen.dart';

class ProfileScreen extends StatelessWidget {
  const ProfileScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Screen(
      children: [
        Center(
          child: Column(
            children: [
              SizedBox(
                width: 200,
                height: 200,
                child: ClipRRect(
                  borderRadius: BorderRadius.circular(100),
                  child: Image.asset(
                    "assets/images/profile.png",
                    fit: BoxFit.cover,
                  ),
                ),
              ),
              const Padding(
                padding: EdgeInsets.symmetric(
                  vertical: 10,
                ),
                child: Text(
                  "Ivan Ivanov",
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 28,
                  ),
                ),
              ),
              const Text(
                "Administrator",
                style: TextStyle(
                  color: CupertinoColors.systemBlue,
                  fontWeight: FontWeight.bold,
                  fontSize: 20,
                ),
              ),
            ],
          ),
        ),
        const Padding(
          padding: EdgeInsets.symmetric(
            vertical: 10,
          ),
          child: Text(
            "Orders",
            style: TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 16,
            ),
          ),
        ),
        OrderListView(
          orders: [
            Order(
              address: "address",
              comment: "comment",
              configurations: [
                Configuration(
                  createdAt: DateTime.now(),
                  id: "id",
                  price: 15000000,
                )
              ],
              createdAt: DateTime.now(),
              customer: Customer(
                id: '',
                email: 'ivan.ivanov@gmail.com',
                phoneNumber: '88005553535',
                firstName: 'Ivan',
                lastName: 'Ivanov',
              ),
              id: "",
              number: 1,
              priceLimit: 15000000,
              status: "In Process",
            )
          ],
        ),
      ],
    );
  }
}
