class Ticket {
  final String id;
  final String title;
  final String description;
  final String userId;
  final String? assignedTo;
  final String status;

  Ticket({
    required this.id,
    required this.title,
    required this.description,
    required this.userId,
    this.assignedTo,
    required this.status,
  });

  factory Ticket.fromJson(Map<String, dynamic> json) => Ticket(
    id: json['id'],
    title: json['title'],
    description: json['description'],
    userId: json['user_id'],
    assignedTo: json['assigned_to'],
    status: json['status'],
  );
}
