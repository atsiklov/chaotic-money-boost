import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';

void main() {
  runApp(const ChaoticMoneyApp());
}

class ChaoticMoneyApp extends StatelessWidget {
  const ChaoticMoneyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.yellow),
        scaffoldBackgroundColor: Colors.black,
      ),
      home: const HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(
          'Chaotic Money',
          style: TextStyle(fontWeight: FontWeight.w900),
        ),
        centerTitle: true,
      ),
      body: ListView.separated(
        itemCount: 50,
        separatorBuilder: (context, i) => Divider(color: Colors.white24),
        itemBuilder: (context, i) => ListTile(
          textColor: Colors.white70,
          leading: SvgPicture.asset('images/tian.svg', height: 50, width: 50),
          title: Text("Challenge ${i + 1}"),
          subtitle: Text(
            "Description for challenge ${i + 1} will be here",
            style: TextStyle(color: const Color.fromARGB(196, 255, 193, 7)),
          ),
          trailing: Icon(Icons.arrow_forward_ios),
        ),
      ),
    );
  }
}
