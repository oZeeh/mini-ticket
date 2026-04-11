import 'package:flutter/material.dart';
import 'package:frontend/screens/login_screen.dart';
import 'package:provider/provider.dart';
import '../../providers/auth_provider.dart';
import '../../providers/ticket_provider.dart';
import '../../widgets/ticket_card.dart';

class UserListScreen extends StatefulWidget {
  const UserListScreen({super.key});

  @override
  State<UserListScreen> createState() => _UserListScreenState();
}

class _UserListScreenState extends State<UserListScreen> {
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      context.read<TicketProvider>().loadAllTickets();
    });
  }

  @override
  Widget build(BuildContext context) {
    final tickets = context.watch<TicketProvider>();
    final auth = context.read<AuthProvider>();

    return Scaffold(
      appBar: AppBar(
        title: const Text('Todos os Tickets'),
        actions: [
          IconButton(
            icon: const Icon(Icons.logout),
            onPressed: () async {
              await auth.logout();
              if (!mounted) return;
              Navigator.pushReplacement(
                context,
                MaterialPageRoute(builder: (_) => const LoginScreen()),
              );
            },
          ),
        ],
      ),
      body: tickets.isLoading
          ? const Center(child: CircularProgressIndicator())
          : tickets.tickets.isEmpty
          ? const Center(child: Text('Nenhum ticket encontrado'))
          : ListView.builder(
              itemCount: tickets.tickets.length,
              itemBuilder: (context, index) {
                final ticket = tickets.tickets[index];
                return TicketCard(ticket: ticket, onTap: () {});
              },
            ),
    );
  }
}
