import 'package:flutter/material.dart';
import 'package:frontend/screens/login_screen.dart';
import 'package:provider/provider.dart';
import '../../providers/auth_provider.dart';
import '../../providers/ticket_provider.dart';
import '../../widgets/ticket_card.dart';
import 'create_ticket_screen.dart';
import 'ticket_detail_screen.dart';

class TicketListScreen extends StatefulWidget {
  const TicketListScreen({super.key});

  @override
  State<TicketListScreen> createState() => _TicketListScreenState();
}

class _TicketListScreenState extends State<TicketListScreen> {
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      final auth = context.read<AuthProvider>();
      context.read<TicketProvider>().loadMyTickets(auth.userId!);
    });
  }

  @override
  Widget build(BuildContext context) {
    final tickets = context.watch<TicketProvider>();
    final auth = context.read<AuthProvider>();

    return Scaffold(
      appBar: AppBar(
        title: const Text('Meus Tickets'),
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
                return TicketCard(
                  ticket: ticket,
                  onTap: () => Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (_) => TicketDetailScreen(ticket: ticket),
                    ),
                  ),
                );
              },
            ),
      floatingActionButton: FloatingActionButton(
        onPressed: () async {
          await Navigator.push(
            context,
            MaterialPageRoute(builder: (_) => const CreateTicketScreen()),
          );
          if (!mounted) return;
          context.read<TicketProvider>().loadMyTickets(auth.userId!);
        },
        child: const Icon(Icons.add),
      ),
    );
  }
}
